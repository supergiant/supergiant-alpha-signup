import { Component, OnInit } from '@angular/core';
import {Validators, FormGroup} from '@angular/forms';
import {FormlyFieldConfig} from 'ng-formly';

import { Http, Response, Headers } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

@Component({
  selector: 'app-request',
  templateUrl: './request.component.html',
  styleUrls: ['./request.component.css']
})
export class RequestComponent implements OnInit {
  user = {
    email: '',
    name: '',
    company: '',
  };

  state = {
    error: false,
    status: 0,
    message: ''
  }

  form: FormGroup = new FormGroup({});
  userFields: FormlyFieldConfig = [{
    fieldGroup: [{
      key: 'email',
      type: 'input',
      templateOptions: {
        type: 'email',
        //label: 'Email address*',
        placeholder: 'Email Address'
      },
      validators: {
        validation: Validators.compose([Validators.required])
      }
    },{
      noFormControl: true,
      template: '<small class="form-text text-muted">We\'ll never share your email with anyone else.</small>'
    }, {
      key: 'name',
      type: 'input',
      templateOptions: {
        type: 'name',
        //label: 'Email address*',
        placeholder: 'Name'
      },
      validators: {
        validation: Validators.compose([Validators.required])
      }
    }, {
      key: 'company',
      type: 'input',
      templateOptions: {
        type: 'company',
        //label: 'Email address*',
        placeholder: 'Company'
      },
      validators: {
        validation: Validators.compose([Validators.required])
      }
    }]
  }];

  getInvite(){
    // http://alpha.supergiant.io/api/request
    this.http.post('http://localhost:8080/request',JSON.stringify({email: this.user.email,name: this.user.name,company: this.user.company}),).map(response => response.json()).subscribe(
      (result) => { if (result["error"]) {
        this.state.status=1
        this.state.error=true;
        this.state.message = result["error"]
      } else {
        this.state.status=1
        this.state.message = result["Thank you, you will receive an invite as soon as a slot opens up"]
      }
    }
    );
  }
  constructor(private http: Http) { }

  ngOnInit() {
  }

}
