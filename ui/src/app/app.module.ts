import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FormlyModule, FormlyBootstrapModule} from 'ng-formly';


import { AppComponent } from './app.component';
import { RequestComponent } from './request/request.component';
import { ClaimComponent } from './claim/claim.component';

@NgModule({
  declarations: [
    AppComponent,
    RequestComponent,
    ClaimComponent
  ],
  imports: [
      NgbModule.forRoot(),
      BrowserModule,
      FormsModule,
      ReactiveFormsModule,
      FormlyModule.forRoot(),
      FormlyBootstrapModule,
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
