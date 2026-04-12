import { Injectable, signal, computed } from '@angular/core';
import { Dish, type Order } from '../models/dish.model';

@Injectable({
  providedIn: 'root',
})
export class BasketService {
  private items = signal<Order[]>([]);

  itemsList = this.items.asReadonly();

  count = computed(() => this.items().length);

  addToBasket(dish: Dish, count: number) {
    this.items.update((old) => [...old, { count: count, dish: dish }]);
  }
}
