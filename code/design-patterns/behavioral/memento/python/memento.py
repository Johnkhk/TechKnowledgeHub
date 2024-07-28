class TextDocument:
    def __init__(self, content=""):
        self._content = content

    def get_content(self):
        return self._content

    def set_content(self, content):
        self._content = content


class Memento:
    def __init__(self, state):
        self._state = state

    def get_state(self):
        return self._state


class TextEditor:
    def __init__(self, document):
        self._document = document
        self._history = []

    def write(self, text):
        self._document.set_content(self._document.get_content() + text)

    def save(self):
        self._history.append(Memento(self._document.get_content()))

    def undo(self):
        if self._history:
            memento = self._history.pop()
            self._document.set_content(memento.get_state())

    def get_content(self):
        return self._document.get_content()
