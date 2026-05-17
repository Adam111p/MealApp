import { Routes } from '@angular/router';
import { MenuList } from './components/menu-list/menu-list';
import { SearchDish } from './search-dish/search-dish';

export const routes: Routes = [
  {
    path: '',
    component: MenuList,
  },
  {
    path: 'search',
    component: SearchDish,
  },
  {
    path: '**',
    redirectTo: '',
  },
];
