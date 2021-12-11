import { Injectable } from '@nestjs/common';
import { CreateTransactionDto } from './dto/create-transaction.dto';
import { PrismaService } from '../prisma.service';

@Injectable()
export class TransactionsService {
  constructor(private prisma: PrismaService) {}

  async create(data: CreateTransactionDto) {
    const transactionCreated = await this.prisma.transaction.create({ data });
    return transactionCreated;
  }

  async findAll() {
    const transactions = await this.prisma.transaction.findMany();
    return transactions;
  }
}
