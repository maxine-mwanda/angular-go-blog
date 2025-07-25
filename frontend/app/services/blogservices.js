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
    },  
    createPost: function(post) {
        return $http.post('/api/posts', post);
      },
    updatePost: function(slug, post) {
        return $http.put('/api/posts/' + slug, post);
      },
    deletePost: function(slug) {
        return $http.delete('/api/posts/' + slug);
      }
};
});