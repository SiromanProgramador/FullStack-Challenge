import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { tap } from 'rxjs/operators';
import { map } from 'rxjs/operators';
@Injectable({
  providedIn: 'root'
})
export class AuthorDetailService {

  constructor(private http: HttpClient) { }

  getOneAuthor(_id: string, headers?:any) {
    return this.http.get<any>('http://localhost:8000/authors/'+ _id, headers)
    .pipe(
        map((data: any) => data)
    );
  }
}
