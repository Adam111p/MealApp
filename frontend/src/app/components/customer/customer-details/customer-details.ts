import { CommonModule } from '@angular/common';
import { Component, signal, ChangeDetectorRef } from '@angular/core';
import { ReactiveFormsModule, Validators, FormBuilder, FormGroup } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AddressForm } from '../address-form/address-form';
import { MatDividerModule } from '@angular/material/divider';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'ma-customer-details',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    AddressForm,
    MatInputModule,
    MatFormFieldModule,
    MatCheckboxModule,
    MatButtonModule,
    MatCardModule,
    AddressForm,
    MatDividerModule,
    MatIconModule,
  ],
  templateUrl: './customer-details.html',
  styleUrl: './customer-details.scss',
})
export class CustomerDetails {
  clientForm: FormGroup;

  specialAddressOrder = signal<boolean>(false);

  constructor(
    private fb: FormBuilder,
    private cdr: ChangeDetectorRef,
  ) {
    this.clientForm = this.fb.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      tel: [''],
      address: this.createAddressGroup(),
      orderAddress: [null],
    });
  }
  createAddressGroup(): FormGroup {
    return this.fb.group({
      street: ['', Validators.required],
      city: ['', Validators.required],
      postCode: ['', Validators.required],
    });
  }
  save() {
    console.log(this.clientForm.value);
  }
  getAddressGroup(name: string) {
    return this.clientForm.get(name) as FormGroup;
  }

  toggleOrderAddress(isChecked: boolean) {
    this.specialAddressOrder.set(isChecked);
    if (isChecked) {
      this.clientForm.setControl('orderAddress', this.createAddressGroup());
    } else {
      this.clientForm.removeControl('orderAddress');
    }
    this.cdr.detectChanges();
  }
}
