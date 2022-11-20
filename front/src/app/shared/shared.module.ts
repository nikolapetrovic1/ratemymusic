import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule,FormsModule  } from '@angular/forms';
import { SharedRoutingModule } from './shared-routing.module';
import {MatFormFieldModule, MatLabel} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { HttpClientModule } from '@angular/common/http';
import { HeaderComponent } from './components/header/header.component';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatCardModule} from '@angular/material/card';
import {NgbRatingModule} from '@ng-bootstrap/ng-bootstrap';
import {MatIconModule} from '@angular/material/icon';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import {MatTableModule} from '@angular/material/table';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
@NgModule({
  declarations: [
    HeaderComponent,
  ],
  imports: [
    CommonModule,
    SharedRoutingModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatButtonModule,
    MatInputModule,
    MatToolbarModule,
    HttpClientModule,
    MatCardModule,
    FormsModule,
    NgbRatingModule,
    MatIconModule,
    MatAutocompleteModule,
    MatSnackBarModule,
    MatTableModule,
    MatProgressBarModule,
    MatProgressSpinnerModule
  ],
  exports: [
    HeaderComponent,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatButtonModule,
    MatInputModule,
    HttpClientModule,
    MatToolbarModule,
    MatCardModule,
    FormsModule,
    NgbRatingModule,
    MatIconModule,
    MatAutocompleteModule,
    MatSnackBarModule,
    MatTableModule,
    MatProgressBarModule,
    MatProgressSpinnerModule
  ]
})
export class SharedModule { }
