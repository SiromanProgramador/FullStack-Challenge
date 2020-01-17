import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthorsComponent } from './pages/authors/authors.component';
import { BooksComponent} from './pages/books/books.component';


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
    path:'**',
    redirectTo: 'books'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
