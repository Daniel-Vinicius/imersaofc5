import { Controller, Get, Post, Body } from '@nestjs/common';
import { TransactionsService } from './transactions.service';
import { CreateTransactionDto } from './dto/create-transaction.dto';

@Controller('transactions')
export class TransactionsController {
  constructor(private readonly transactionsService: TransactionsService) {}

  @Post()
  async create(@Body() body: CreateTransactionDto) {
    const { amount, account_id } = body;

    const transactionCreated = await this.transactionsService.create({
      amount,
      account_id,
    });

    return transactionCreated;
  }

  @Get()
  async findAll() {
    const transactions = await this.transactionsService.findAll();
    return transactions;
  }
}
