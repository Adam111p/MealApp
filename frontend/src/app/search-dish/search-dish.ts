import { Component, inject, signal } from '@angular/core';
import {
  FormGroup,
  FormBuilder,
  Validators,
  ɵInternalFormsSharedModule,
  ReactiveFormsModule,
} from '@angular/forms';

import { MatFormField } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { LoadingStatus, MenuService } from '../services/menu.service';
import type { Dish } from '../models/dish.model';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatIconModule } from '@angular/material/icon';
import { OrderDetails } from '../components/order-details/order-details';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatButtonModule } from '@angular/material/button';


@Component({
  selector: 'ma-search-dish',
  imports: [
    MatFormField,
    MatInputModule,
    ɵInternalFormsSharedModule,
    ReactiveFormsModule,
    MatCardModule,
    MatChipsModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatButtonModule,
    MatDialogModule,
  ],
  templateUrl: './search-dish.html',
  styleUrl: './search-dish.scss',
})
export class SearchDish {
protected readonly LoadingStatus = LoadingStatus;
  [x: string]: any;
  search() {
    this.menuService.searchDishes(this.clientForm.controls['descriptionQuery'].value || '');
  }
  loadingStatus = LoadingStatus.Loading;
  clientForm: FormGroup;
  query = signal<string>('');

  menuService = inject(MenuService);
  private dialog = inject(MatDialog);

  constructor(private fb: FormBuilder) {
    this.clientForm = this.fb.group({
      descriptionQuery: ['', Validators.minLength(3)],
    });
  }

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
