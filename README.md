# go_test_starter

A demo for golang TDD

relate blog: https://williaminfante.medium.com/golang-testing-with-tdd-e548d8be776

TDD stands for Test-Driven Development.

Development Cycle:
1. Write a test before develop. The test should define the desired behavior and functionality of the change.
1. Run this test. Will get a failed answer, but don't worry, the change has not been inplemented yet.
1. Write the code. Finish your feature by the simpal change. Do not need to strive for optimization and quality.The focus should on making the test pass.
1. Run all the test. After you finish your change, you should keep all the test should be pass at this point to guarantee functional integrity and correctness.
1. Refactor. At this point, your code is usually not very clear and maintainable. You should refactor the code to improve its design, structure and perfoemance.
1. Rerun all the test. You should make sure your logic has not change in previous step.
1. Repeat. You can repeat this process when you get a new feature.
