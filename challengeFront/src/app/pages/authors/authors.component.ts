import { Component, OnInit } from '@angular/core';
import { AuthorService } from 'src/app/shared/author.service';

@Component({
  selector: 'app-authors',
  templateUrl: './authors.component.html',
  styleUrls: ['./authors.component.scss']
})
export class AuthorsComponent implements OnInit {

  authors = [];
  ruta = 
    {
      name: 'Author-Detail',
      path: '/author-detail'
    }
  

  constructor(
    private authorService: AuthorService
  ) {  }

  ngOnInit() {
    this.authorService.getAllAuthors().subscribe((response) => {
      this.authors = response;
    });
  }

}
