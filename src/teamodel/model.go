package teamodel

import (
	"crypto/rsa"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"

	"github.com/m1kkY8/lockbox/src/comps"
	"github.com/m1kkY8/lockbox/src/config"
	"github.com/m1kkY8/lockbox/src/encryption"
	"github.com/m1kkY8/lockbox/src/styles"
)

type Model struct {
	// User info
	username    string
	userColor   string
	currentRoom string
	// Model comps
	width           int
	height          int
	input           textinput.Model
	viewport        comps.Model
	onlineUsers     comps.Model
	styles          *styles.Styles
	messageList     *MessageList
	messageChannel  chan string
	onlineUsersChan chan []string
	PublicKeysChan  chan []*rsa.PublicKey
	// Server side
	conn       *websocket.Conn
	keyPair    *encryption.RSAKeys
	PublicKeys []*rsa.PublicKey
}

type MessageList struct {
	messages []string
	count    int
}

var messageLimit = 100

func New(conf config.Config, conn *websocket.Conn, keyPair *encryption.RSAKeys) *Model {
	styles := styles.DefaultStyle(conf.Color)

	input := textinput.New()
	input.Prompt = ""
	input.Placeholder = "Join any room to start typing"
	input.Width = 50
	input.Focus()

	vp := comps.New(50, 20)
	vp.SetContent("Welcome, start messaging")

	onlineList := comps.New(20, 20)
	onlineList.SetContent("Online")

	return &Model{
		conn:            conn,
		userColor:       conf.Color,
		username:        conf.Username,
		input:           input,
		styles:          styles,
		viewport:        vp,
		onlineUsers:     onlineList,
		messageList:     &MessageList{},
		messageChannel:  make(chan string),
		onlineUsersChan: make(chan []string),
		PublicKeys:      []*rsa.PublicKey{},
		currentRoom:     "",
		keyPair:         keyPair,
	}
}

func (m *Model) View() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				m.styles.Border.Render(m.viewport.View()),
				m.styles.Border.Render(m.onlineUsers.View()),
			),
			m.styles.Border.Render(m.input.View()),
		),
	)
}

func (m *Model) Init() tea.Cmd {
	go m.recieveMessages()
	return tea.Batch(
		m.listenForMessages(),
		m.listenForOnlineUsers(),
	)
}
