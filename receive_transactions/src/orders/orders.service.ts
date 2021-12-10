import { EmptyResultError } from 'sequelize';
import { Inject, Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { AccountStorageService } from 'src/accounts/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';
import { Producer } from '@nestjs/microservices/external/kafka.interface';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order)
    private orderModel: typeof Order,
    private accountStorage: AccountStorageService,
    @Inject('KAFKA_PRODUCER')
    private kafkaProducer: Producer,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const {
      credit_card_number,
      credit_card_cvv,
      credit_card_expiration_month,
      credit_card_expiration_year,
    } = createOrderDto;

    const accountId = this.accountStorage.account.id;

    const orderCreated = await this.orderModel.create({
      account_id: accountId,
      ...createOrderDto,
    });

    await this.kafkaProducer.send({
      topic: 'transactions',
      messages: [
        {
          key: accountId,
          value: JSON.stringify({
            id: orderCreated.id,
            account_id: accountId,
            credit_card_number,
            credit_card_expiration_year,
            credit_card_expiration_month,
            credit_card_cvv,
            amount: orderCreated.amount,
          }),
        },
      ],
    });

    return orderCreated;
  }

  async findAll() {
    const orders = await this.orderModel.findAll({
      where: {
        account_id: this.accountStorage.account.id,
      },
    });

    return orders;
  }

  async findOne(id: string) {
    const order = await this.orderModel.findOne({
      where: {
        id: id,
        account_id: this.accountStorage.account.id,
      },
      rejectOnEmpty: new EmptyResultError(`Order with ID ${id} not found`),
    });

    return order;
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const order = await this.orderModel.findByPk(id);
    await order.update(updateOrderDto);
  }
}
