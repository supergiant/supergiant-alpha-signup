import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {FormlyModule, FormlyBootstrapModule} from 'ng-formly';
import { RouterModule, Routes } from '@angular/router';
import { HttpModule } from '@angular/http';
import { Angulartics2Module, Angulartics2GoogleTagManager } from 'angulartics2';

import { AppComponent } from './app.component';
import { RequestComponent } from './request/request.component';
import { ClaimComponent } from './claim/claim.component';
import { ClosedComponent } from './closed/closed.component';

const appRoutes: Routes = [
  // { path: 'request', component: RequestComponent },
  // { path: 'claim',      component: ClaimComponent },
  { path: '', component: ClosedComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    RequestComponent,
    ClaimComponent,
    ClosedComponent
  ],
  imports: [
      NgbModule.forRoot(),
      BrowserModule,
      FormsModule,
      HttpModule,
      ReactiveFormsModule,
      FormlyModule.forRoot(),
      FormlyBootstrapModule,
      Angulartics2Module.forRoot([ Angulartics2GoogleTagManager ]),
      RouterModule.forRoot(
      appRoutes
    )
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
