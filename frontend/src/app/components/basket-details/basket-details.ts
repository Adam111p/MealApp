import { Component } from '@angular/core';
import { CustomerDetails } from '../customer/customer-details/customer-details';

@Component({
  selector: 'ma-basket-details',
  imports: [CustomerDetails],
  templateUrl: './basket-details.html',
  styleUrl: './basket-details.scss',
})
export class BasketDetails {}
