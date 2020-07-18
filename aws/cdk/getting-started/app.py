#!/usr/bin/env python3

from aws_cdk import core

from getting_started.getting_started_stack import GettingStartedStack


app = core.App()
GettingStartedStack(app, "getting-started")

app.synth()
