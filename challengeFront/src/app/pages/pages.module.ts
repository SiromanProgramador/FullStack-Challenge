import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthorsComponent } from './authors/authors.component';
import { BooksComponent } from './books/books.component';



@NgModule({
  declarations: [
    AuthorsComponent,
    BooksComponent
  ],
  exports: [
    AuthorsComponent,
    BooksComponent

  ],
  imports: [
    CommonModule
  ]
})
export class PagesModule { }
