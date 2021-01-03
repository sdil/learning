defmodule LearnPlugTest do
  use ExUnit.Case
  doctest LearnPlug

  test "greets the world" do
    assert LearnPlug.hello() == :world
  end
end
