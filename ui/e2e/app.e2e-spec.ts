import { SupergiantAlphaSignupPage } from './app.po';

describe('supergiant-alpha-signup App', () => {
  let page: SupergiantAlphaSignupPage;

  beforeEach(() => {
    page = new SupergiantAlphaSignupPage();
  });

  it('should display welcome message', done => {
    page.navigateTo();
    page.getParagraphText()
      .then(msg => expect(msg).toEqual('Welcome to app!!'))
      .then(done, done.fail);
  });
});
