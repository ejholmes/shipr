module Shipr
  module Hooks
    class GitHub < Grape::API
      logger Shipr.logger

      format :json
      default_format :json

      helpers do
        def event
          ActiveSupport::StringInquirer.new(headers['X-Github-Event'])
        end

        def deploy
          GitHubJobCreator.create(params)
        end
      end

      helpers do
        delegate :authenticate!, to: :warden

        def warden; env['warden'] end
      end

      before do
        authenticate!(scope: :api)
      end

      params do
        optional :id, type: Integer
        optional :sha, type: String
        optional :name, type: String
        optional :description, type: String
        group :payload do
          optional :force, type: Boolean
          optional :environment, type: String
          optional :config, type: Hash
        end
      end
      post do
        status 200
        if event.deployment?
          present deploy
        elsif event.deployment_status?
          Notifiers::Slack.notify(params)
        else
          {}
        end
      end
    end
  end
end