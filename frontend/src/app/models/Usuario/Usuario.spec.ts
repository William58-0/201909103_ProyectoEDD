import { Usuario } from './Usuario';

describe('Usuario', () => {
  it('should create an instance', () => {
    expect(new Usuario(0,"","","","")).toBeTruthy();
  });
});