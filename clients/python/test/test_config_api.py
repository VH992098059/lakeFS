"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""


import unittest

import lakefs_client
from lakefs_client.api.config_api import ConfigApi  # noqa: E501


class TestConfigApi(unittest.TestCase):
    """ConfigApi unit test stubs"""

    def setUp(self):
        self.api = ConfigApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_config(self):
        """Test case for get_config

        """
        pass

    def test_get_garbage_collection_config(self):
        """Test case for get_garbage_collection_config

        """
        pass


if __name__ == '__main__':
    unittest.main()
