import { Controller, Get, Post, Body, UseGuards } from '@nestjs/common';
import { OrdersService } from './orders.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { TokenGuard } from 'src/accounts/token.guard';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { KafkaMessage } from '@nestjs/microservices/external/kafka.interface';
import { OrderStatus } from './entities/order.entity';

@UseGuards(TokenGuard)
@Controller('orders')
export class OrdersController {
  constructor(private readonly ordersService: OrdersService) {}

  @Post()
  async create(@Body() createOrderDto: CreateOrderDto) {
    const order = await this.ordersService.create(createOrderDto);
    return order;
  }

  @Get()
  async findAll() {
    const orders = await this.ordersService.findAll();
    return orders;
  }

  @MessagePattern('transactions_result')
  async consumerUpdateStatus(@Payload() message: KafkaMessage) {
    const data = message.value as any;
    const { id, status } = data as { id: string; status: OrderStatus };
    await this.ordersService.update(id, { status });
  }
}
