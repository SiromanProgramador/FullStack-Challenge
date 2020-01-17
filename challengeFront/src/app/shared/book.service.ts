import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { tap } from 'rxjs/operators';
import { map } from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class BookService {

  constructor( private http: HttpClient ) { }


    getPosts() {
        return this.http.get('localhost:8000')
            .pipe( tap( console.log ) );
    }

    updateBook(multipart: any, headers?: any) {
        return this.http.patch<any>('http://localhost:8000/books', multipart, headers)
        .pipe(
            map((data: any) => data)
        );
    }
    createBook(multipart: any, headers?: any) {
        return this.http.post<any>('http://localhost:8000/books', multipart, headers)
        .pipe(
            map((data: any) => data)
        );
    }
    getAllBooks(headers?: any) {
        return this.http.get<any>('http://localhost:8000/books',headers)
        .pipe(
            map((data: any) => data)
        );

    }
}