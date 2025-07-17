angular.module('angular-go-blog.services', []).factory('BlogService', function($http) {
  return {
    getPosts: function() {
      return $http.get('/api/posts').then(function(response) {
        return response.data;
      });
    },
    getPost: function(slug) {
      return $http.get('/api/posts/' + slug).then(function(response) {
        return response.data;
      });
    }
  };
});