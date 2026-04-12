import { Component, inject, signal } from '@angular/core';
import {
  MAT_DIALOG_DATA as DIALOG_DATA,
  MatDialogModule,
  MatDialogRef as REF,
} from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { Dish, type Spice } from '../../models/dish.model';
import { BasketService } from '../../services/basket.service';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MenuService } from '../../services/menu.service';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'ma-order-details',
  styleUrl: './order-details.scss',
  imports: [MatDialogModule, MatButtonModule, MatCheckboxModule, MatIconModule],
  templateUrl: './order-details.html',
})
export class OrderDetails {
  [x: string]: any;
  changeQuantity(addCount: number) {
    this.quantity.set(Math.max(this.quantity() + addCount, 1));
  }
  data: Dish = inject(DIALOG_DATA);
  dialogRef = inject(REF<OrderDetails>);
  private basketService = inject(BasketService);

  menuService = inject(MenuService);
  quantity = signal<number>(1);
  selectedSpices = new Map<Spice, boolean>();

  isChecked(spice: Spice) {
    return !!this.selectedSpices.get(spice);
  }

  toggleSpice(spice: Spice) {
    if (this.selectedSpices.get(spice)) {
      this.selectedSpices.set(spice, false);
    } else this.selectedSpices.set(spice, true);
  }
  confirm() {
    this.dialogRef.close(this.data);
    this.data.spices = [...this.selectedSpices.entries()]
      .filter(([spice, selected]) => selected)
      .map(([spice, selected]) => spice);
    this.basketService.addToBasket(this.data, this.quantity());
  }
}
