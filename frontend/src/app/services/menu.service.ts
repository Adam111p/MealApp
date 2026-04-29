import { inject, Injectable, signal } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { toSignal } from '@angular/core/rxjs-interop';
import { Dish, type Spice } from '../models/dish.model';

@Injectable({
  providedIn: 'root',
})
export class MenuService {
  private http = inject(HttpClient);
  private apiUrl: string = 'http://localhost:8080/api';

  // toSignal automatycznie subskrybuje i zamienia wynik na Signal
  // initialValue zapewnia, że nie dostaniemy błędu przy starcie
  menuItems = toSignal(this.http.get<Dish[]>(this.apiUrl + '/menu'), { initialValue: [] });
  spicesDict = signal<Spice[]>([]);
  dishes = signal<Dish[]>([]);
  loadStatus = signal<LoadingStatus>(LoadingStatus.OK);
  getSpicesDictByType(type: string) {
    return this.spicesDict().filter((s) => s.typeDish === type);
  }
  loadSpices() {
    this.http.get<Spice[]>(this.apiUrl + '/spices').subscribe((data) => {
      this.spicesDict.set(data);
    });
  }

  searchDishes(query: string) {
    this.loadStatus.set(LoadingStatus.Loading);

    this.http.get<Dish[]>(this.apiUrl + '/searchDishes?query=' + query).subscribe({
      next: (data) => {
        this.dishes.set(data);
        this.loadStatus.set(LoadingStatus.OK);
      },
      error: (err) => {
        this.loadStatus.set(LoadingStatus.Error);
      },
    });
  }
}
export enum LoadingStatus {
  OK,
  Loading,
  Error,
}
