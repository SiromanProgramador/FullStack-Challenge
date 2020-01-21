import { Component, OnInit } from '@angular/core';
import { AuthorDetailService } from 'src/app/shared/author-detail.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-author-detail',
  templateUrl: './author-detail.component.html',
  styleUrls: ['./author-detail.component.scss']
})
export class AuthorDetailComponent implements OnInit {

  author;
  authorId: string

  constructor(
    private authorDetailService: AuthorDetailService,
    private activateRoute: ActivatedRoute
  ) { 
    this.authorId = this.activateRoute.snapshot.params['id'];
  }

  ngOnInit() {

   this.authorDetailService.getOneAuthor(this.authorId).subscribe((response) => {
     this.author = response;
   })
  }

}
