import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class BookDetailService {

  constructor( private http: HttpClient) { }

  getOneBook(_id:string, headers?:any) {
    return this.http.get<any>('http://localhost:8000/books/'+_id, headers)
    .pipe(
      map((data: any) => data)
    );
  }
}
