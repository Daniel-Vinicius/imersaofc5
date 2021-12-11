import { HttpException, HttpStatus, Injectable } from '@nestjs/common';
import { CreateTransactionDto } from './dto/create-transaction.dto';
import { PrismaService } from '../prisma.service';

@Injectable()
export class TransactionsService {
  constructor(private prisma: PrismaService) {}

  async create(data: CreateTransactionDto) {
    const countTransactionsByAccount = await this.prisma.transaction.count({
      where: {
        account_id: data.account_id,
      },
    });

    if (countTransactionsByAccount >= 3) {
      throw new HttpException(
        'Limit of transactions by account reached',
        HttpStatus.BAD_REQUEST,
      );
    }

    const transactionCreated = await this.prisma.transaction.create({ data });
    return transactionCreated;
  }

  async findAll() {
    const transactions = await this.prisma.transaction.findMany();
    return transactions;
  }
}
