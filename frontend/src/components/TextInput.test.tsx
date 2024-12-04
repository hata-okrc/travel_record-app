// TextInput.test.js
import userEvent from "@testing-library/user-event";
import { render, screen } from "@testing-library/react";
import TextInput from "./textInput";

test("TextInput Component Test", async () => {
  const user = userEvent.setup();
  render(<TextInput />);

  const textElement = screen.getByLabelText("Entered Text");
  expect(textElement);

  const inputElement = screen.getByLabelText("Text Input");
  await user.type(inputElement, "How beautiful");
  expect(screen.getByText("Entered Text: Hello World"));
});
