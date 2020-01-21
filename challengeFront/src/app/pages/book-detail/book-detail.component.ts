import { Component, OnInit } from '@angular/core';
import { BookDetailService } from 'src/app/shared/book-detail.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-book-detail',
  templateUrl: './book-detail.component.html',
  styleUrls: ['./book-detail.component.scss']
})
export class BookDetailComponent implements OnInit {

  book;
  bookId: string

  constructor(
    private bookDetailService: BookDetailService,
    private activateRoute: ActivatedRoute
  ) { 
    this.bookId = this.activateRoute.snapshot.params['id']
  }

  ngOnInit() {
    
    this.bookDetailService.getOneBook(this.bookId).subscribe((response) => {
      this.book = response
    })
  }

}
