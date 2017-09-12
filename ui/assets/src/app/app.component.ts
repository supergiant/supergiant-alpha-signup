import { Angulartics2GoogleTagManager } from 'angulartics2';

import { Component } from '@angular/core';

import {Validators, FormGroup} from '@angular/forms';
import {FormlyFieldConfig} from 'ng-formly';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'app'
  constructor(angulartics2GoogleTagManager: Angulartics2GoogleTagManager) {}
}
