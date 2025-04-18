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

from lakefs_sdk.models.stats_events_list import StatsEventsList  # noqa: E501

class TestStatsEventsList(unittest.TestCase):
    """StatsEventsList unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> StatsEventsList:
        """Test StatsEventsList
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `StatsEventsList`
        """
        model = StatsEventsList()  # noqa: E501
        if include_optional:
            return StatsEventsList(
                events = [
                    lakefs_sdk.models.stats_event.StatsEvent(
                        class = '', 
                        name = '', 
                        count = 56, )
                    ]
            )
        else:
            return StatsEventsList(
                events = [
                    lakefs_sdk.models.stats_event.StatsEvent(
                        class = '', 
                        name = '', 
                        count = 56, )
                    ],
        )
        """

    def testStatsEventsList(self):
        """Test StatsEventsList"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
