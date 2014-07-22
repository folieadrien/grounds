class GroundsController < ApplicationController
  def show
    @ground = GroundDecorator.new(Ground.new(language: 'golang'), view_context)
  end

  def run
    grounder = Grounder.new
    stdout, stderr = grounder.exec(params[:ground][:language], params[:ground][:code])
    respond_to do |format|
      format.js { render json: { stdout: stdout.join, stderr: stderr.join, status: :ok } }
    end
  end
  
  def switch_editor_theme
    session[:editor_theme] = params[:editor_theme]
  end
end
