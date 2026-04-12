export interface Dish {
  id: number;
  name: string;
  price: number;
  description: string;
  fileName: string;
  typeDish: string;
  spices: Spice[];
}
export interface Spice {
  id: number;
  name: string;
  levelSpice: number;
  description: string;
  typeDish: string;
}

export interface Customer {
  firstName: string;
  lastName: string;
}
export interface Basket {
  customer?: Customer;
  order: Order;
}

export interface Order {
  dish: Dish;
  count: number;
}
