import { Component, Input } from '@angular/core';
import { ReactiveFormsModule, type FormGroup } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'ma-address-form',
  imports: [MatInputModule, ReactiveFormsModule],
  templateUrl: './address-form.html',
  styleUrl: './address-form.scss',
})
export class AddressForm {
  @Input() addressGroup!: FormGroup;
}
