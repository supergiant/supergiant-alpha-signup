import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FormlyModule, FormlyBootstrapModule} from 'ng-formly';
import { RouterModule, Routes } from '@angular/router';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { RequestComponent } from './request/request.component';
import { ClaimComponent } from './claim/claim.component';

const appRoutes: Routes = [
  { path: 'request', component: RequestComponent },
  { path: '',      component: ClaimComponent },
];

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
      HttpModule,
      ReactiveFormsModule,
      FormlyModule.forRoot(),
      FormlyBootstrapModule,
      RouterModule.forRoot(
      appRoutes
    )
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
