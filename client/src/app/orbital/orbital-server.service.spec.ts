import { TestBed } from '@angular/core/testing';

import { OrbitalServerService } from './orbital-server.service';

describe('OrbitalServerService', () => {
  let service: OrbitalServerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OrbitalServerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
