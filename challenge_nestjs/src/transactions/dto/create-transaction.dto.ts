import { IsUUID, IsNotEmpty, IsNumber, IsPositive } from 'class-validator';

export class CreateTransactionDto {
  @IsNotEmpty()
  @IsNumber()
  @IsPositive()
  amount: number;

  @IsNotEmpty()
  @IsUUID()
  account_id: string;
}
