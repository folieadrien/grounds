class GroundsController < ApplicationController
  def show
    ground = Ground.new(language: session[:language] || GroundEditor.default_option_code(:language),
                        code: '')
    @ground = GroundDecorator.new(ground, view_context)
  end

  def shared
    ground = Ground.from_storage(params[:id])
    @ground = GroundDecorator.new(ground, view_context)
    render 'show'
  end

  def share
    @ground = Ground.new(ground_params)
    @ground.save
    render json: { status: :ok, shared_url: grounds_shared_url(@ground.id) }
  end
  
  def switch_option
    option, value = params[:option], params[:value]
    if option.present? && value.present? && GroundEditor.has_option?(option, value)
      session[option] = value
    end
    render json: { status: :ok }
  end

  private

  def ground_params
    params.require(:ground)
  end
end
