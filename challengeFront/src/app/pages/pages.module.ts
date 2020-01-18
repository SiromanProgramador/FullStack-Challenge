import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AuthorsComponent } from './authors/authors.component';
import { BooksComponent } from './books/books.component';
import { AuthorDetailComponent } from './author-detail/author-detail.component';
import { BookDetailComponent } from './book-detail/book-detail.component';



@NgModule({
  declarations: [
    AuthorsComponent,
    BooksComponent,
    AuthorDetailComponent,
    BookDetailComponent
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
