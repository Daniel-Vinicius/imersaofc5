export class CreateOrderDto {
  amount: number;
  credit_card_number: number;
  credit_card_name: string;
  credit_card_cvv: number;
  credit_card_expiration_month: number;
  credit_card_expiration_year: number;
}
