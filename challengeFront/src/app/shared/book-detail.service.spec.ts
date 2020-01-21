import { TestBed } from '@angular/core/testing';

import { BookDetailService } from './book-detail.service';

describe('BookDetailService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: BookDetailService = TestBed.get(BookDetailService);
    expect(service).toBeTruthy();
  });
});
