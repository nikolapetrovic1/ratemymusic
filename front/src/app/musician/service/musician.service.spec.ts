import { TestBed } from '@angular/core/testing';

import { MusicianService } from './musician.service';

describe('MusicianService', () => {
  let service: MusicianService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MusicianService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
