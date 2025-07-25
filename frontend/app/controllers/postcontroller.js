angular.module('angular-go-blog.controllers').controller('PostController', 
  function($scope, $routeParams, BlogService) {
    $scope.loading = true;
    $scope.post = {};
    
    BlogService.getPost($routeParams.slug).then(function(post) {
      $scope.post = post;
      $scope.loading = false;
    }).catch(function(error) {
      console.error('Error fetching post:', error);
      $scope.loading = false;
    });
 
    // Initialize
    $scope.loading = true;
    $scope.isEditing = false;
    $scope.post = {};
    $scope.newPost = { title: '', content: '', excerpt: '' };

    // Fetch post if slug exists (edit/view mode)
    if ($routeParams.slug) {
      BlogService.getPost($routeParams.slug)
        .then(function(post) {
          $scope.post = post;
          $scope.loading = false;
        })
        .catch(function(error) {
          console.error('Error fetching post:', error);
          toastr.error('Failed to load post');
          $scope.loading = false;
        });
    } else {
      // New post mode
      $scope.isEditing = true;
      $scope.loading = false;
    }

    // Create a new post
    $scope.createPost = function() {
      $scope.loading = true;
      BlogService.createPost($scope.newPost)
        .then(function(response) {
          toastr.success('Post created successfully!');
          $location.path('/posts/' + response.slug);
        })
        .catch(function(error) {
          console.error('Error creating post:', error);
          toastr.error('Failed to create post');
        })
        .finally(function() {
          $scope.loading = false;
        });
    };

    // Update existing post
    $scope.updatePost = function() {
      $scope.loading = true;
      BlogService.updatePost($scope.post.slug, $scope.post)
        .then(function(response) {
          toastr.success('Post updated successfully!');
          $scope.isEditing = false;
        })
        .catch(function(error) {
          console.error('Error updating post:', error);
          toastr.error('Failed to update post');
        })
        .finally(function() {
          $scope.loading = false;
        });
    };

    // Delete post
    $scope.deletePost = function() {
      if (!confirm('Are you sure you want to delete this post?')) return;
      
      $scope.loading = true;
      BlogService.deletePost($scope.post.slug)
        .then(function() {
          toastr.success('Post deleted successfully!');
          $location.path('/');
        })
        .catch(function(error) {
          console.error('Error deleting post:', error);
          toastr.error('Failed to delete post');
        })
        .finally(function() {
          $scope.loading = false;
        });
    };

    // Toggle edit mode
    $scope.toggleEdit = function() {
      $scope.isEditing = !$scope.isEditing;
    };
    }
);