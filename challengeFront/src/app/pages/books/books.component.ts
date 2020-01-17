import { Component, OnInit } from '@angular/core';
import { BookService } from 'src/app/shared/book.service';

@Component({
  selector: 'app-books',
  templateUrl: './books.component.html',
  styleUrls: ['./books.component.scss']
})
export class BooksComponent implements OnInit {

  books = [];

  constructor(
    private bookService: BookService
  ) { }

  ngOnInit() {
    this.bookService.getAllBooks().subscribe((response) => {
      this.books = response;
    })
  }

}
