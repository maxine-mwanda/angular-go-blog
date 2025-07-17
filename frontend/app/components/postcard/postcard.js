angular.module('angular-go-blog.components', []).component('postCard', {
  bindings: {
    post: '<'
  },
  templateUrl: 'app/components/post-card/post-card.html'
});