from aws_cdk import (
        core,
        aws_s3_assets as assets
)


class GettingStartedStack(core.Stack):

    def __init__(self, scope: core.Construct, id: str, **kwargs) -> None:
        super().__init__(scope, id, **kwargs)

        # The code that defines your stack goes here
        # bucket = s3.Bucket(self, "sitebucket", bucket_name="fadhil-getting-started-bucket", public_read_access=True, website_index_document="index.html")
        # core.CfnOutput(self, "sitebucketname", value=bucket.bucket_name)
        # core.CfnOutput(self, "siteBucketWebsite", value=bucket.bucket_website_url)

        asset = assets.Asset(self, "SampleAsset", path="./sample-asset/index.html")

