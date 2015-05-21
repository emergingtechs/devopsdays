set :css_dir, 'stylesheets'

set :js_dir, 'javascripts'

set :images_dir, 'images'

activate :directory_indexes
activate :bootstrap_navbar
activate :sprockets

# Reload the browser automatically whenever files change
configure :development do
  activate :livereload
end


# Build-specific configuration
configure :build do
  activate :asset_hash
end

activate :deploy do |deploy|
  deploy.build_before = true
  deploy.method = :git
  deploy.branch = "gh-pages"
end
