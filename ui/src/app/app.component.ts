import { Component } from '@angular/core';

import {Validators, FormGroup} from '@angular/forms';
import {FormlyFieldConfig} from 'ng-formly';



@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'app';

    user = {
      email: '',
      invite: '',
    };


  form: FormGroup = new FormGroup({});
   userFields: FormlyFieldConfig = [{
     className: 'row',
     fieldGroup: [{
         className: 'col-md-2',
         key: 'email',
         type: 'input',
         templateOptions: {
             type: 'email',
             label: 'Email address',
             placeholder: 'Enter email'
         },
         validators: {
           validation: Validators.compose([Validators.required])
         }
     }, {
         className: 'col-md-2',
         key: 'invite',
         type: 'input',
         templateOptions: {
             type: 'string',
             label: 'Invite Code',
             placeholder: 'INVITECODE',
             pattern: ''
         },
         validators: {
           validation: Validators.compose([Validators.required])
         }
     }]
   }];

  submit(user) {
    console.log(user);
  }
}
