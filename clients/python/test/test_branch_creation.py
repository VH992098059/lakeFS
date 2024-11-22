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

import lakefs_sdk
from lakefs_sdk.models.branch_creation import BranchCreation  # noqa: E501
from lakefs_sdk.rest import ApiException

class TestBranchCreation(unittest.TestCase):
    """BranchCreation unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test BranchCreation
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `BranchCreation`
        """
        model = lakefs_sdk.models.branch_creation.BranchCreation()  # noqa: E501
        if include_optional :
            return BranchCreation(
                name = '', 
                source = '', 
                force = True, 
                hidden = True
            )
        else :
            return BranchCreation(
                name = '',
                source = '',
        )
        """

    def testBranchCreation(self):
        """Test BranchCreation"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
