import { Component, inject, signal, type OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenuList } from './components/menu-list/menu-list';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatBadgeModule } from '@angular/material/badge';
import { BasketService } from './services/basket.service';
import { MenuService } from './services/menu.service';
import { MatDialog } from '@angular/material/dialog';
import { BasketDetails } from './components/basket-details/basket-details';
@Component({
  selector: 'ma-root',
  imports: [
    RouterOutlet,
    MenuList,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatBadgeModule,
  ],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App implements OnInit {
  private menuService = inject(MenuService);
  private dialog = inject(MatDialog);
  ngOnInit(): void {
    this.menuService.loadSpices();
  }
  protected readonly title = signal('frontend');
  basketService = inject(BasketService);
  showBasket() {
    this.dialog.open(BasketDetails, {
      height: '800px',
      minWidth: '700px',
    });
  }
}
