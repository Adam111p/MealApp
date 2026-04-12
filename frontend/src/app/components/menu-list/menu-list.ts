import { Component, inject } from '@angular/core';
import { MenuService } from '../../services/menu.service';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatChipsModule } from '@angular/material/chips';
import { MatDialog } from '@angular/material/dialog';
import { Dish } from '../../models/dish.model';
import { OrderDetails } from '../order-details/order-details';

@Component({
  selector: 'ma-menu-list',
  imports: [MatCardModule, MatButtonModule, MatIconModule, MatChipsModule, MatIconModule],
  standalone: true,
  templateUrl: './menu-list.html',
  styleUrl: './menu-list.scss',
})
export class MenuList {
  menuService = inject(MenuService);
  private dialog = inject(MatDialog);

  openDetails(dish: Dish) {
    const dialogRef = this.dialog.open(OrderDetails, {
      data: dish,
      width: '600px',
      maxWidth: '90vw',
      panelClass: 'custom-dialog-container',
      enterAnimationDuration: '300ms',
      exitAnimationDuration: '200ms',
    });
    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        console.log('Dodano do koszyka:', result);
      }
    });
  }
}
