import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthorsComponent } from './pages/authors/authors.component';
import { BooksComponent} from './pages/books/books.component';
import { BookDetailComponent } from './pages/book-detail/book-detail.component';
import { AuthorDetailComponent } from './pages/author-detail/author-detail.component';


const routes: Routes = [
  {
    path: 'authors',
    component: AuthorsComponent
  },
  {
    path: 'books',
    component: BooksComponent
  },
  {
    path: 'book-detail/:id',
    component: BookDetailComponent
  },
  {
    path: 'author-detail/:id',
    component: AuthorDetailComponent
  },
  {
    path:'**',
    redirectTo: 'books'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
