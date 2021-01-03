defmodule Example.Application do
  use Application
  require Logger

  def start(_type, _args) do
    children = [
      {Plug.Cowboy, scheme: :http, plug: Example.Router, options: [port: 8080]}]
      opts = [strategy: :one_for_one, name: Example.Supervisor]

      Logger.info("Starting applications")

      Supervisor.start_link(children, opts)
  end
end
