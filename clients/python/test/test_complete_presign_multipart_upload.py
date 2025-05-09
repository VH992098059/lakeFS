# coding: utf-8

"""
    lakeFS API

    lakeFS HTTP API

    The version of the OpenAPI document: 1.0.0
    Contact: services@treeverse.io
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from lakefs_sdk.models.complete_presign_multipart_upload import CompletePresignMultipartUpload  # noqa: E501

class TestCompletePresignMultipartUpload(unittest.TestCase):
    """CompletePresignMultipartUpload unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CompletePresignMultipartUpload:
        """Test CompletePresignMultipartUpload
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CompletePresignMultipartUpload`
        """
        model = CompletePresignMultipartUpload()  # noqa: E501
        if include_optional:
            return CompletePresignMultipartUpload(
                physical_address = '',
                parts = [
                    lakefs_sdk.models.upload_part.UploadPart(
                        part_number = 56, 
                        etag = '', )
                    ],
                user_metadata = {
                    'key' : ''
                    },
                content_type = ''
            )
        else:
            return CompletePresignMultipartUpload(
                physical_address = '',
                parts = [
                    lakefs_sdk.models.upload_part.UploadPart(
                        part_number = 56, 
                        etag = '', )
                    ],
        )
        """

    def testCompletePresignMultipartUpload(self):
        """Test CompletePresignMultipartUpload"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
