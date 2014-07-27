#!/usr/bin/env rake

namespace :images do
  registry = ARGV[1] || ''

  abort("Please specifiy a docker registry!") if registry.empty?

  task :build do
    images do |file, path|
      sh "docker build -t #{registry}/#{file} #{path}"
    end
  end

  task :push do
    images do |file, path|
      sh "docker push #{registry}/#{file}"
    end
  end

  task :pull do
    images do |file, _|
      sh "docker pull #{registry}/#{file}"
    end
  end
end

def images
  return unless block_given?

  dockerfiles = File.join("#{File.dirname(__FILE__)}", 'dockerfiles')

  Dir.entries(dockerfiles).select { |f| f != '.' && f != '..' }.each do |file|
    path = File.join(dockerfiles, file)
    if File.directory?(path)
      yield(file, path)
    end
  end
end